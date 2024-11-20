using Microsoft.AspNetCore.Mvc;
using Server.Services;
using System.Collections.Generic;
using System.Linq;

namespace Server.Controllers
{
	[Route("api/[controller]")]
	[ApiController]
	public class AuthController : ControllerBase
	{
		// Храним пользователей в списке
		private static List<LoginModel> _users = new List<LoginModel>
		{
			new LoginModel { Username = "user1", Password = "password1" },
			new LoginModel { Username = "user2", Password = "password2" },
			new LoginModel { Username = "admin", Password = "adminpass" }
		};

		// Метод для авторизации и получения JWT токена
		[HttpPost("login")]
		public ActionResult Login([FromBody] LoginModel model)
		{
			// Ищем пользователя в списке
			var user = _users.FirstOrDefault(u => u.Username == model.Username && u.Password == model.Password);

			if (user == null)
			{
				return Unauthorized();  // Если пользователь не найден или пароль неверен
			}

			var token = JwtService.GenerateToken(model.Username);  // Генерируем JWT токен
			return Ok(new { token });
		}

		// Получить всех пользователей
		[HttpGet("users")]
		public ActionResult GetAllUsers()
		{
			return Ok(_users);
		}

		// Получить пользователя по имени
		[HttpGet("user/{username}")]
		public ActionResult GetUser(string username)
		{
			var user = _users.FirstOrDefault(u => u.Username == username);
			if (user == null)
				return NotFound();

			return Ok(user);
		}

		// Добавить нового пользователя
		[HttpPost("user")]
		public ActionResult AddUser([FromBody] LoginModel user)
		{
			// Проверяем, если пользователь с таким именем уже существует
			if (_users.Any(u => u.Username == user.Username))
				return BadRequest("Пользователь с таким именем уже существует.");

			_users.Add(user);
			return CreatedAtAction(nameof(GetUser), new { username = user.Username }, user);
		}

		// Обновить данные пользователя
		[HttpPut("user/{username}")]
		public ActionResult UpdateUser(string username, [FromBody] LoginModel updatedUser)
		{
			var user = _users.FirstOrDefault(u => u.Username == username);
			if (user == null)
				return NotFound();

			user.Username = updatedUser.Username;
			user.Password = updatedUser.Password;

			return NoContent();
		}

		// Удалить пользователя
		[HttpDelete("user/{username}")]
		public ActionResult DeleteUser(string username)
		{
			var user = _users.FirstOrDefault(u => u.Username == username);
			if (user == null)
				return NotFound();

			_users.Remove(user);
			return NoContent();
		}
	}

	// Модель для авторизации и данных пользователя
	public class LoginModel
	{
		public string Username { get; set; }
		public string Password { get; set; }
	}
}
