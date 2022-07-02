import java.util.Scanner;

public class EntryPoint 
{
	public static void main(String[] args) 
	{
//		CalculatorGUI gui = new CalculatorGUI();
//		gui.Show();
		Scanner scanner = new Scanner(System.in);
		var str = scanner.nextLine();
		for (var s : str.split("[-*/+]"))
		{
			System.out.println("Input: " + s);	
		}
		
		scanner.close();
	}
}
