using System;

namespace Relax.MmoGame.Common
{
    public interface IPackage<out TCommand> : IPackageSerializable where TCommand : struct, IConvertible
    {
        TCommand Command { get; }
    }
}