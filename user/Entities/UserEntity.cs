namespace User.Entities;

public class UserEntity
{
    public int Id { get; set; }
    public string Login { get; set; } = null!;
    public string Email { get; set; } = null!;
    public string Password { get; set; } = null!;
}