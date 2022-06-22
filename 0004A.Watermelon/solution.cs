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
            int w = Convert.ToInt32(System.Console.ReadLine());
            System.Console.WriteLine("{0}", ((w > 2 && w % 2 == 0) ? "YES" : "NO"));
        }

        static void Main(string[] args)
        {
            Solution();
        }
    }
}
