import java.awt.BorderLayout;
import java.awt.GridLayout;
import java.awt.event.*;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedHashSet;
import java.util.Set;

import javax.swing.*;

public class CalculatorGUI 
{
	private JFrame m_frame = new JFrame();
	private JPanel m_digitPanel = new JPanel();
	private JPanel m_operationPanel = new JPanel();
	private JLabel m_label = new JLabel();
	private StringBuffer m_buffer = new StringBuffer();
	private ArrayList<JButton> m_digitButtons = new ArrayList<JButton>();
	private ArrayList<JButton> m_operationButtons = new ArrayList<JButton>();
	private boolean m_isOpInputPossible = false;
	
	private final LinkedHashSet<String> OPERATIONS = 
			new LinkedHashSet<String>(Arrays.asList("+", "-", "*", "/", "="));
	private final LinkedHashSet<String> DIGITS = 
			new LinkedHashSet<String>(Arrays.asList("0", "1", "2", "3", "4", "5", "6", "7", "8", "9"));
	
	private void SetPanels(JPanel currentPanel, Set<String> buttonValues, ButtonTypes type)
	{
		for (String buttonValue : buttonValues)
		{
			JButton button = new JButton(buttonValue);
			switch (type)
			{
				case OPERATION:
				{
					button.setEnabled(false);
					m_operationButtons.add(button);
					break;
				}
				case DIGIT:
				{
					button.setEnabled(true);
					m_digitButtons.add(button);
					break;
				}
			}
			button.addActionListener(new ButtonListener());
			button.setSize(10, 10);
			currentPanel.add(button);
		}
	}
	
	private void ProcessButtonEnable()
	{
		for (JButton button : m_operationButtons)
		{
			button.setEnabled(m_isOpInputPossible);
		}
	}
	
	public CalculatorGUI()
	{	
		SetPanels(m_digitPanel, DIGITS, ButtonTypes.DIGIT);
		SetPanels(m_operationPanel, OPERATIONS, ButtonTypes.OPERATION);
		
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
			if (DIGITS.contains(e.getActionCommand()))
			{
				m_isOpInputPossible = true;	
			}
			else
			{
				m_isOpInputPossible = false;
			}
			if (!e.getActionCommand().equals("="))
			{
				m_buffer.append(e.getActionCommand());
			}
			else
			{
				
				m_buffer.setLength(0);
			}
			m_label.setText(m_buffer.toString());	
			ProcessButtonEnable();
		}
	}
	
}
