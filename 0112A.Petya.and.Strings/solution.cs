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
            String first = System.Console.ReadLine().ToUpper();
            String second = System.Console.ReadLine().ToUpper();

            int result = 0;
            for (int i = 0; i < first.Length; ++i)
            {
                if (first[i] < second[i])
                {
                    result = -1;
                    break;
                }

                if (first[i] > second[i])
                {
                    result = 1;
                    break;
                }
            }

            System.Console.WriteLine(result);
        }

        static void Main(string[] args)
        {
            Solution();
        }
    }
}
