using System.Runtime.CompilerServices;
using Microsoft.EntityFrameworkCore;
using User.Entities;
using User.Infrastructure.Persistence;

namespace User.Infrastructure.Seeding;

public static class HostExtensions
{
    public static void SeedDatabase(this IHost host)
    {
        using (var scope = host.Services.CreateScope())
        {
            var services = scope.ServiceProvider;

            var context = services.GetRequiredService<UserDbContext>();
            if (context.Database.GetPendingMigrations().Any())
            {
                context.Database.Migrate();
            }

            SeedRoles(context);
        }
    }

    private static void SeedRoles(UserDbContext context)
    {
        UserRoleEntity[] roles = new[]
        {
            new UserRoleEntity { Name =  "Teacher" },
            new UserRoleEntity { Name =  "Student" },
        };
        SaveList(roles, context);
    }

    private static void SaveList<T>(IEnumerable<T> entities, UserDbContext context) where T : class
    {
        if (!context.Set<T>().Any())
        {
            context.AddRange(entities);
            context.SaveChanges();
        }
    }
}