namespace Relax.MmoGame.Common
{
    public static class Consts
    {
        public const int Fps = 2;
        
        // MTU 1500 bytes

        public const int MaxClientRead = 1000;
        public const int MaxClientWrite = 50;
        
        public const int MaxServerRead = MaxClientWrite;
        public const int MaxServerWrite = MaxClientRead;
    }
}