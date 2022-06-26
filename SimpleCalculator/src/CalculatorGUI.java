import java.awt.BorderLayout;
import java.awt.GridLayout;
import java.awt.event.*;

import javax.swing.*;

public class CalculatorGUI 
{
	private JFrame m_frame;
	private JPanel m_digitPanel;
	private JPanel m_operationPanel;
	private JLabel m_label;
	
	private void InitFields()
	{
		m_frame = new JFrame();
		m_digitPanel = new JPanel();
		m_operationPanel = new JPanel();
		m_label = new JLabel();
	}
	
	private void SetPanels(JPanel currentPanel, char buttonValues[])
	{
		for (char buttonValue : buttonValues)
		{
			JButton button = new JButton(String.valueOf(buttonValue));
			button.addActionListener(new ButtonListener());
			button.setSize(10, 10);
			currentPanel.add(button);
		}
	}
	
	public CalculatorGUI()
	{
		InitFields();

		char digits[] = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'};
		char operations[] = {'+', '-', '*', '/', '='};
		
		SetPanels(m_digitPanel, digits);
		SetPanels(m_operationPanel, operations);
		
		m_digitPanel.setLayout(new GridLayout(2, 5));
		m_operationPanel.setLayout(new GridLayout(1, 5));
		m_label.setText("<Your input>");
		
		m_frame.getContentPane().add(BorderLayout.CENTER, m_digitPanel);
		m_frame.getContentPane().add(BorderLayout.SOUTH, m_operationPanel);
		m_frame.getContentPane().add(BorderLayout.NORTH, m_label);
		m_frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		m_frame.setSize(300, 300);
	}
	
	public void Show()
	{
		m_frame.setVisible(true);
	}
	
	
	class ButtonListener implements ActionListener
	{
		public void actionPerformed(ActionEvent e) 
		{
			if (!e.getActionCommand().equals("="))
			{
				m_label.setText(e.getActionCommand());	
			}
		}
	}
	
}
