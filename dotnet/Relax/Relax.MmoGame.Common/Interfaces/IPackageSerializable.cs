using System;

namespace Relax.MmoGame.Common
{
    public interface IPackageSerializable
    {
        // todo calculate size by reflection
        int Size { get; }

        byte[] Serialize(Span<byte> buffer);
    }
}