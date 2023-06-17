using Microsoft.EntityFrameworkCore;
using User.Entities;

namespace User.Infrastructure.Persistence;

public class UserDbContext : DbContext
{
    public UserDbContext() { }
    public UserDbContext(DbContextOptions<UserDbContext> opt) : base(opt) { }
    public DbSet<UserEntity> Users { get; set; } = null!;

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        base.OnModelCreating(modelBuilder);

        modelBuilder.Entity<UserEntity>(user =>
        {
            user.HasKey(e => e.Id);
            user.HasIndex(e => e.Login).IsUnique();
            user.HasIndex(e => e.Email).IsUnique();
        });
    }
}