import java.util.ArrayList;

public class Calculator 
{
	static public int GetResultFromString(String inputString)
	{
		var nodeList = GetNodesFromString(inputString);
		int result = 0;
		for (var node : nodeList)
		{
			result += node.GetResult();
		}
		return result;
	}
	
	static public ArrayList<Node> GetNodesFromString(String inputString)
	{
		ArrayList<Node> result = new ArrayList<Node>();
		
		String[] s = inputString.toString().split("[+-*/]");
		
		return result;
	}
	
}
