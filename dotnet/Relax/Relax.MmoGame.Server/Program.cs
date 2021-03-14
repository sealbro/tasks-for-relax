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
            // const int worldSize = 10;
            // const int playersCount = 5;
            //
            // var map = new WorldMap(worldSize);
            // var playerWatcher = new PlayerWatcher(worldSize);
            //
            // var players = new RegisteredPlayer[playersCount];
            // for (int i = 0; i < playersCount; i++)
            // {
            //     players[i] = playerWatcher.RegisterPlayer();
            // }
            //
            // while (true)
            // {
            //     map.Draw(players);
            //
            //     foreach (var player in players)
            //     {
            //         var moveDirection = BotPlayer.GetDirection();
            //         Console.WriteLine(moveDirection);
            //         playerWatcher.MovePlayer(player, moveDirection);
            //     }
            //
            //     await Task.Delay(1000);
            // }


            var builder = new HostBuilder()
                .ConfigureServices((_, services) =>
                {
                    services
                        .AddHostedService<ServerHostService>();
                });

            await builder.RunConsoleAsync();
        }
    }
}