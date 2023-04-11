using System;
using System.Collections.Generic;
using System.Linq;
using System.Runtime.InteropServices;
using System.Threading;
using Relax.Dotnet.System.Collections;

namespace Relax.Dotnet
{
    class Program
    {
        static void Main(string[] args)
        {
        }

        private static string GetAddr( object obj)
        {
            GCHandle gch = GCHandle.Alloc(obj, GCHandleType.Pinned);
            IntPtr pObj = gch.AddrOfPinnedObject();
            return pObj.ToString("X");
        }
    }
}