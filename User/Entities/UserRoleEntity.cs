namespace User.Entities;

public class UserRoleEntity
{
    public int Id { get; set; }
    public string Name { get; set; } = null!;
    public List<UserEntity> Users { get; set; } = new();
}