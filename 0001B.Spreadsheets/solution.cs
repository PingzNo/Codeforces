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
            int caseCount = int.Parse(Console.ReadLine());
            for (int caseIndex = 0; caseIndex < caseCount; caseIndex++)
            {
                string text = Console.ReadLine();
                int pos = text.IndexOf('C');

                if (text[0] == 'R' && Char.IsDigit(text[1]) && 1 < pos && pos < text.Length)
                {
                    // R23C55 -> BC23
                    int row = Convert.ToInt32(text.Substring(1, pos - 1));
                    int col = Convert.ToInt32(text.Substring(pos + 1));

                    List<char> listText = new List<char>();
                    while (col > 0)
                    {
                        if (col % 26 == 0)
                        {
                            listText.Add('Z');
                            col = (col - 1) / 26;
                        } else {
                            listText.Add((char)('A' + (col % 26) - 1));
                            col /= 26;
                        }
                    }

                    while (listText.Count > 0)
                    {
                        Console.Write(listText[listText.Count - 1]);
                        listText.RemoveAt(listText.Count - 1);
                    }

                    Console.WriteLine(row);
                }
                else
                {
                    // BC23 -> R23C55
                    int i = 0;
                    for (; i < text.Length; i++)
                    {
                        if (Char.IsDigit(text[i]))
                        {
                            break;
                        }
                    }

                    string row = text.Substring(0, i);
                    string col = text.Substring(i);

                    int accum = 0;
                    foreach (char r in row)
                    {
                        accum = accum * 26 + (int)(r - 'A') + 1;
                    }

                    Console.WriteLine("R{0}C{1}", col, accum);
                }
            }
        }

        static void Main(string[] args)
        {
            Solution();
        }
    }
}
