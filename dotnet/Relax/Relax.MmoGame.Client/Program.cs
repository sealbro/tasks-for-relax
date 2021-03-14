using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

namespace Relax.MmoGame.Client
{
    class Program
    {
        static async Task Main(string[] args)
        {
            var builder = new HostBuilder()
                .ConfigureServices((_, services) =>
                {
                    services
                        .AddHostedService<ClientHostService>();
                });

            await builder.RunConsoleAsync();
        }
    }
}