FROM mcr.microsoft.com/dotnet/sdk:7.0-bullseye-slim AS build
WORKDIR /source

# copy csproj and restore as distinct layers
ADD . .
RUN dotnet restore

# copy everything else and build app
RUN dotnet publish -c release -o /app --no-restore

# final stage/image
FROM mcr.microsoft.com/dotnet/aspnet:7.0-bullseye-slim
RUN apt update && apt upgrade -y && apt install curl -y
WORKDIR /app
COPY --from=build /app ./
ENTRYPOINT ["dotnet", "User.dll"]
