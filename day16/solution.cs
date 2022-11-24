class Solution
{
    public static void Main(string[] args)
    {
        string S = Console.ReadLine();
        
        try
        {
            var i = int.Parse(S);
            Console.Write(i);
            
        }
        catch
        {
            Console.Write("Bad String");
        }
    }
}
