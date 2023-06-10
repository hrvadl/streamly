using Grpc.Core;
using User;
using User.Infrastructure.Persistence;

namespace User.Services;

public class UsersService : Users.UsersBase
{
    private readonly ILogger<UsersService> logger;
    private readonly UserDbContext context;

    public UsersService(ILogger<UsersService> logger, UserDbContext context)
    {
        this.logger = logger;
        this.context = context;
    }

    public override Task<GetByIdResponse> GetById(GetByIdRequest request, ServerCallContext context)
    {
        return Task.FromResult(new GetByIdResponse
        {
            Id = "guid",
            Name = "Name",
            Email = "Email.com",
            Number = "+380500",
            Description = null,
            Role = new UserRole
            {
                Id = 1,
                Name = "Student",
            },
        });
    }
}
