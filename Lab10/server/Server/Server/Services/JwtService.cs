using Microsoft.IdentityModel.Tokens;
using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Text;

namespace Server.Services
{
	public class JwtService
	{
		private const string SecretKey = "supawdawdawderseawcretIOPAafwionfawpifoihoawinconoaiwkey12f34f5"; // Для простоты
		private const string Issuer = "Server";
		private const string Audience = "user";

		public static string GenerateToken(string username)
		{
			var securityKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(SecretKey));
			var credentials = new SigningCredentials(securityKey, SecurityAlgorithms.HmacSha256);
			var claims = new List<System.Security.Claims.Claim>
			{
				new System.Security.Claims.Claim(System.Security.Claims.ClaimTypes.Name, username)
			};

			var token = new JwtSecurityToken(Issuer, Audience, claims, expires: DateTime.Now.AddMinutes(30), signingCredentials: credentials);
			return new JwtSecurityTokenHandler().WriteToken(token);
		}

		public static bool ValidateToken(string token)
		{
			var securityKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(SecretKey));
			var tokenHandler = new JwtSecurityTokenHandler();

			try
			{
				tokenHandler.ValidateToken(token, new Microsoft.IdentityModel.Tokens.TokenValidationParameters
				{
					ValidateIssuer = true,
					ValidateAudience = true,
					ValidateLifetime = true,
					IssuerSigningKey = securityKey,
					ValidIssuer = Issuer,
					ValidAudience = Audience
				}, out var validatedToken);
				return true;
			}
			catch
			{
				return false;
			}
		}
	}
}
