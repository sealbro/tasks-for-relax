using System;
using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Relax.MmoGame.Common;

namespace Relax.MmoGame.Server
{
    class Program
    {
        static async Task Main(string[] args)
        {
            var map = new WorldMap();
            map.Draw(new PlayerPosition {Id = 1}, new PlayerPosition {Id = 2, X = 4},
                new PlayerPosition {Id = 3, Y = 6}, new PlayerPosition {Id = 4, X = 9, Y = 9});

            // var builder = new HostBuilder()
            //     .ConfigureServices((_, services) =>
            //     {
            //         services
            //             .AddHostedService<ServerHostService>();
            //     });
            //
            // await builder.RunConsoleAsync();
        }
    }
}