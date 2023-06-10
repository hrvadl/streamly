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

        modelBuilder.Entity<UserEntity>()
            .HasKey(e => e.Id);
    }
}