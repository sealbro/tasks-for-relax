using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Hosting;

namespace Relax.MmoGame.Server
{
    public class ServerHostService: IHostedService
    {
        public Task StartAsync(CancellationToken cancellationToken)
        {
            throw new System.NotImplementedException();
        }

        public Task StopAsync(CancellationToken cancellationToken)
        {
            throw new System.NotImplementedException();
        }
    }
}