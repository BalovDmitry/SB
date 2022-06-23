import javax.swing.*;

import java.awt.BorderLayout;
import java.awt.event.*;

public class SimpleGui
{
	public SimpleGui()
	{
		m_panel = new MyDrawPanel()	;	
		m_frame = new JFrame();
		m_label = new JLabel();
		m_frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		JButton colorButton = new JButton("repaint");
		colorButton.addActionListener(new ButtonListener());
		
		JButton labelButton = new JButton("change description");
		labelButton.addActionListener(new LabelListener());
		
		m_frame.getContentPane().add(BorderLayout.SOUTH, colorButton);
		m_frame.getContentPane().add(BorderLayout.CENTER, m_panel);
		
		m_frame.getContentPane().add(BorderLayout.EAST, labelButton);
		m_frame.getContentPane().add(BorderLayout.WEST, m_label);
		
		m_frame.setSize(300, 300);
	}
	
	public void Show()
	{
		m_frame.setVisible(true);
	}
	
	class LabelListener implements ActionListener
	{
		public void actionPerformed(ActionEvent e) {
			m_label.setText("Hello world!");
			
		}
	}
	
	class ButtonListener implements ActionListener
	{
		public void actionPerformed(ActionEvent e) {
			m_frame.repaint();
			
		}
	}
	
	private JFrame m_frame;
	private MyDrawPanel m_panel;
	private JLabel m_label;
}
