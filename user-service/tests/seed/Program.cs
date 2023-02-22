using Bogus;
using Microsoft.EntityFrameworkCore;

Randomizer.Seed = new Random(8675309);

var dbServer = Environment.GetEnvironmentVariable("DB_SERVER");
var dbPort = Environment.GetEnvironmentVariable("DB_PORT");
var dbName = Environment.GetEnvironmentVariable("DB_NAME");
var dbUsername = Environment.GetEnvironmentVariable("DB_USER");
var dbPassword = Environment.GetEnvironmentVariable("DB_PASS");

Console.WriteLine($"Connecting to database {dbName} on server {dbServer},{dbPort} as user {dbUsername}");

string connectionString = $"Server={dbServer},{dbPort};Database={dbName};User={dbUsername};Password={dbPassword};MultipleActiveResultSets=true;TrustServerCertificate=true";

var options = new DbContextOptionsBuilder<ApplicationDbContext>();
options.UseSqlServer(connectionString);

using (var context = new ApplicationDbContext(options.Options))
{
    context.Database.EnsureCreated();

    if (context.Users.Any())
    {
        Console.WriteLine("Table is not empty");
        return;
    }

    var users = new Faker<User>()
        .RuleFor(x => x.Firstname, f => f.Person.FirstName);

    context.Users.AddRange(users.GenerateBetween(400_000, 500_000));

    context.SaveChanges();

    Console.WriteLine("Seed complete!");
}