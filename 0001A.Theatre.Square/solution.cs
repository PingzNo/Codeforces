using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;


namespace Codeforces
{
    class Program
    {
        static void Solution()
        {
            List<UInt64> line = System.Console.ReadLine().Split().Select(UInt64.Parse).ToList();
            UInt64 height = line[0], width = line[1], size = line[2];

            UInt64 height_ratio = (height % size == 0) ? (height / size) : ((height / size) + 1);
            UInt64 width_ratio = (width % size == 0) ? (width / size) : ((width / size) + 1);

            System.Console.WriteLine("{0}", height_ratio * width_ratio);
        }

        static void Main(string[] args)
        {
            Solution();
        }
    }
}
