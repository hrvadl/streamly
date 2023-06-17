using Grpc.Core;
using Microsoft.EntityFrameworkCore;
using User;
using User.Entities;
using User.Infrastructure.Persistence;

namespace User.Services;

public class UsersService : Users.UsersBase
{
    private readonly ILogger<UsersService> logger;
    private readonly UserDbContext dbContext;

    public UsersService(ILogger<UsersService> logger, UserDbContext dbContext)
    {
        this.logger = logger;
        this.dbContext = dbContext;
    }

    public override async Task<CreateResponse> Create(CreateRequest request, ServerCallContext context)
    {
        var user = new UserEntity
        {
            Login = request.Login,
            Email = request.Email,
            Password = request.Password,
        };

        try
        {
            dbContext.Add(user);
            await dbContext.SaveChangesAsync(context.CancellationToken);
        }
        catch (DbUpdateException ex)
        {
            var duplication = await dbContext.Users.Where(u => u.Login == request.Login || u.Email == request.Email).AnyAsync(context.CancellationToken);
            if (duplication)
            {
                throw new RpcException(new Status(StatusCode.InvalidArgument, $"User with such email or login already exists."));
            }
            throw new RpcException(new Status(StatusCode.Internal, $"Exception on user saving occurred.", ex));
        }
        catch (Exception ex)
        {
            throw new RpcException(new Status(StatusCode.Internal, $"Exception on user saving occurred.", ex));
        }

        return new CreateResponse
        {
            Id = user.Id,
        };
    }

    public override async Task<GetByIdResponse> GetById(GetByIdRequest request, ServerCallContext context)
    {
        var user = await dbContext.Users.AsNoTracking().FirstOrDefaultAsync(u => u.Id == request.Id, context.CancellationToken);
        if (user == null)
        {
            throw new RpcException(new Status(StatusCode.InvalidArgument, $"User with id {request.Id} does not exist."));
        }

        return new GetByIdResponse
        {
            Id = user.Id,
            Login = user.Login,
            Email = user.Email,
        };
    }
}
