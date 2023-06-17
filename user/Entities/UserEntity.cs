namespace User.Entities;

public class UserEntity
{
    public int Id { get; set; }
    public string Name { get; set; } = null!;
    public string Email { get; set; } = null!;
    public string Number { get; set; } = null!;
    public string? Description { get; set; }
    public string Password { get; set; } = null!;
    public int RoleId { get; set; }
    public UserRoleEntity Role { get; set; } = null!;
}