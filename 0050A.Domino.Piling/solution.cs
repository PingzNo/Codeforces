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
            List<int> line = System.Console.ReadLine().Split().Select(int.Parse).ToList();
            int height = line[0], width = line[1];

            int half_height = height / 2;
            int half_width = width / 2;
            int result = 2 * half_height * half_width;
            if ((height % 2 == 1) && (width % 2 == 1)) {
                result += half_height + half_width;
            } else if ((height % 2 == 1) && (width % 2 == 0)) {
                result += half_width;
            } else if ((height % 2 == 0) && (width % 2 == 1)) {
                result += half_height;
            }

            System.Console.WriteLine("{0}", result);
        }

        static void Main(string[] args)
        {
            Solution();
        }
    }
}
