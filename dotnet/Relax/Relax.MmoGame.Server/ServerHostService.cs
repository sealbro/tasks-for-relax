using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Hosting;

namespace Relax.MmoGame.Server
{
    public class ServerHostService : IHostedService
    {
        readonly GameServer gameServer = new(24242);

        public async Task StartAsync(CancellationToken cancellationToken)
        {
            await gameServer.Start(cancellationToken);
        }

        public async Task StopAsync(CancellationToken cancellationToken)
        {
            await gameServer.DisposeAsync();
        }
    }
}