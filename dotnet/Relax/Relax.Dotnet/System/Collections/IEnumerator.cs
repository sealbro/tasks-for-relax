using System;
using System.Collections.Generic;

namespace Relax.Dotnet.System.Collections
{
    public interface IEnumerator<out T>
    {
        T Current { get; }
        bool MoveNext();
    }

    public class GuidEnumerator : IEnumerator<string>
    {
        private int _state = -1;

        public string Current { get; private set; }

        public bool MoveNext()
        {
            switch (_state)
            {
                case -1:
                    _state = 0;
                    Current = "state_one_" + Guid.NewGuid();
                    return true;
                case 0:
                    _state = 1;
                    Current = "state_two_" + Guid.NewGuid();
                    return true;
                default:
                    while (true)
                    {
                        _state++;
                        Current = "while_true_" + Guid.NewGuid();
                        return true;
                    }
            }
        }

        public IEnumerable<string> YieldTrue()
        {
            yield return "state_one_" + Guid.NewGuid();
            yield return "state_two_" + Guid.NewGuid();

            while (true)
            {
                yield return "while_true_" + Guid.NewGuid();
            }
        }
    }

    public class TakeEnumerator<T> : IEnumerator<T>
    {
        private readonly IEnumerator<T> _enumerator;
        private int _count;

        public TakeEnumerator(IEnumerator<T> enumerator, int count)
        {
            _enumerator = enumerator;
            _count = count;
        }

        public T Current { get; private set; }

        public bool MoveNext()
        {
            var moveNext = _count > 0 && _enumerator.MoveNext();

            _count--;
            Current = _enumerator.Current;

            return moveNext;
        }
    }

    public static class EnumeratorExtensions
    {
        public static IEnumerator<string> GenerateGuids()
        {
            return new GuidEnumerator();
        }

        public static IEnumerator<T> Take<T>(this IEnumerator<T> enumerator, int count)
        {
            return new TakeEnumerator<T>(enumerator, count);
        }

        public static T[] ToArray<T>(this IEnumerator<T> enumerator)
        {
            var result = Array.Empty<T>();
            var index = -1;

            while (enumerator.MoveNext())
            {
                index++;
                if (result.Length == 0 || result.Length <= index)
                {
                    var buffer = new T[result.Length == 0 ? 2 : result.Length * 2];
                    Array.Copy(result, buffer, index);
                    result = buffer;
                }

                result[index] = enumerator.Current;
            }

            Array.Resize(ref result, index + 1);

            return result;
        }
    }
}