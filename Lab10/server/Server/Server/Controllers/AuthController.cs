using Microsoft.AspNetCore.Mvc;
using Server.Services;

namespace Server.Controllers
{
	[Route("api/[controller]")]
	[ApiController]
	public class AuthController : ControllerBase
	{
		// Метод для авторизации и получения JWT токена
		[HttpPost("login")]
		public ActionResult Login([FromBody] LoginModel model)
		{
			if (model.Username == "user" && model.Password == "password") // Простейшая проверка
			{
				var token = JwtService.GenerateToken(model.Username);
				return Ok(new { token });
			}
			return Unauthorized();
		}
	}

	public class LoginModel
	{
		public string Username { get; set; }
		public string Password { get; set; }
	}
}
