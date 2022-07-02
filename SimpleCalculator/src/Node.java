
public class Node 
{
	private NodeValue m_lhs;
	private NodeValue m_rhs;
	private Operations m_operation;
	
	public int GetResult()
	{
		switch (m_operation)
		{
		case ADD:
			return m_lhs.GetValue() + m_rhs.GetValue();
		case SUBTRACT:
			return m_lhs.GetValue() - m_rhs.GetValue();
		case MULTIPLY:
			return m_lhs.GetValue() * m_rhs.GetValue();
		case DIVIDE:
			return m_lhs.GetValue() / m_rhs.GetValue();
		default:
			return 0;
		}
	}
}
