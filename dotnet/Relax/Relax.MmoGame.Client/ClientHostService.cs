using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Hosting;

namespace Relax.MmoGame.Client
{
    public class ClientHostService : IHostedService
    {
        private readonly GameClient _gameClient = new("127.0.0.1", 24242);

        public async Task StartAsync(CancellationToken cancellationToken)
        {
            await _gameClient.Start(cancellationToken);
        }

        public async Task StopAsync(CancellationToken cancellationToken)
        {
            await _gameClient.DisposeAsync();
        }
    }
}