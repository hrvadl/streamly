using Confluent.Kafka;
using Microsoft.EntityFrameworkCore;
using User.Infrastructure.Persistence;
using User.Infrastructure.Seeding;
using User.Services;

var builder = WebApplication.CreateBuilder(args);

var configuration = builder.Configuration;

// Additional configuration is required to successfully run gRPC on macOS.
// For instructions on how to configure Kestrel and gRPC clients on macOS, visit https://go.microsoft.com/fwlink/?linkid=2099682

// Add services to the container.
builder.Services.AddGrpc();
builder.Services.AddGrpcReflection();

// Kafka
builder.Services.Configure<ProducerConfig>(configuration.GetSection("Kafka"));

builder.Services.AddDbContext<UserDbContext>(options =>
{
    var connectionString = configuration.GetConnectionString("DbConnectionString")!;
    options.UseMySql(connectionString, ServerVersion.AutoDetect(connectionString));
});

var app = builder.Build();

// Configure the HTTP request pipeline.
app.MapGrpcService<UsersService>();
app.MapGrpcReflectionService();
app.MapGet("/", () => "Communication with gRPC endpoints must be made through a gRPC client. To learn how to create a client, visit: https://go.microsoft.com/fwlink/?linkid=2086909");

app.SeedDatabase();

app.Run();
