using Microsoft.AspNetCore.Mvc;
using Server.Models;
using Server.Services;
using System.Collections.Generic;
using System.Linq;
using Server.Models;

namespace Server.Controllers
{
	[Route("api/[controller]")]
	[ApiController]
	public class UsersController : ControllerBase
	{
		private static List<User> users = new List<User>
		{
			new User { Id = 1, Name = "John", Age = 28 },
			new User { Id = 2, Name = "Jane", Age = 32 }
		};

		// Метод для получения всех пользователей
		[HttpGet]
		public ActionResult<IEnumerable<User>> Get()
		{
			return Ok(users);
		}

		// Метод для получения пользователя по ID
		[HttpGet("{id}")]
		public ActionResult<User> Get(int id)
		{
			var user = users.FirstOrDefault(u => u.Id == id);
			if (user == null)
				return NotFound();

			return Ok(user);
		}

		// Метод для добавления нового пользователя
		[HttpPost]
		public ActionResult<User> Create(User user)
		{
			user.Id = users.Max(u => u.Id) + 1;
			users.Add(user);
			return CreatedAtAction(nameof(Get), new { id = user.Id }, user);
		}

		// Метод для обновления данных пользователя
		[HttpPut("{id}")]
		public ActionResult Update(int id, User user)
		{
			var existingUser = users.FirstOrDefault(u => u.Id == id);
			if (existingUser == null)
				return NotFound();

			existingUser.Name = user.Name;
			existingUser.Age = user.Age;
			return NoContent();
		}

		// Метод для удаления пользователя
		[HttpDelete("{id}")]
		public ActionResult Delete(int id)
		{
			var user = users.FirstOrDefault(u => u.Id == id);
			if (user == null)
				return NotFound();

			users.Remove(user);
			return NoContent();
		}
	}
}
