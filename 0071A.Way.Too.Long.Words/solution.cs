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
            int caseCount = Convert.ToInt32(System.Console.ReadLine());
            for (int caseIndex = 0; caseIndex < caseCount; ++caseIndex)
            {
                String word = System.Console.ReadLine();
                if (word.Length <= 10)
                {
                    System.Console.WriteLine(word);
                }
                else
                {
                    System.Console.WriteLine("{0}{1}{2}", word[0], word.Length - 2, word[word.Length - 1]);
                }
            }
        }

        static void Main(string[] args)
        {
            Solution();
        }
    }
}
