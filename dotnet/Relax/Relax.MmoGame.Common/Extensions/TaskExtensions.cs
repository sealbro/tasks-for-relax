using System.Threading;
using System.Threading.Tasks;

namespace Relax.MmoGame.Common.Extensions
{
    public static class TaskExtensions
    {
        public static async void RunBackground(this Task task, CancellationToken cancellationToken)
        {
            await Task.Run(async () => { await task; }, cancellationToken);
        }
    }
}