import java.awt.BorderLayout;
import java.awt.event.*;
import java.util.ArrayList;

import javax.swing.*;

public class CalculatorGUI 
{
	private JFrame m_frame;
	private JTextField m_textField;
	private JPanel m_panel;
	private JLabel m_label;
	private ArrayList<JButton> m_buttons;
	
	private void InitFields()
	{
		m_frame = new JFrame();
		m_textField = new JTextField();
		m_panel = new JPanel();
		m_label = new JLabel();
		m_buttons = new ArrayList<JButton>();
		m_label.setText("Hello world!");
		m_frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
	}
	
	public CalculatorGUI()
	{
		InitFields();
		
		m_frame.getContentPane().add(BorderLayout.CENTER, m_panel);
		m_frame.getContentPane().add(BorderLayout.NORTH, m_textField);
		m_frame.getContentPane().add(BorderLayout.WEST, m_label);
		
		for (int i = 0; i < 10; ++i)
		{
			m_buttons.add(new JButton(String.valueOf(i)));
			m_buttons.get(i).addActionListener(new ButtonListener(i));
			m_frame.getContentPane().add(BorderLayout.CENTER, m_buttons.get(i));
		}
		
		m_frame.setSize(300, 300);
	}
	
	public void Show()
	{
		m_frame.setVisible(true);
	}
	
	
	class ButtonListener implements ActionListener
	{
		public ButtonListener(int number)
		{
			m_number = number;
		}
		
		public void actionPerformed(ActionEvent e) 
		{
			m_label.setText(String.valueOf(m_number));
		}
		
		private int m_number;
	}
	
}
