import java.awt.*;
import javax.swing.*;

public class MyDrawPanel extends JPanel 
{
	public void paintComponent(Graphics g)
	{
		g.setColor(Color.orange);
		g.fillRect(0, 0, this.getWidth(), this.getHeight());
		int red = (int)(Math.random() * 255);
		int green = (int)(Math.random() * 255);
		int blue = (int)(Math.random() * 255);
		
		Color randomColor = new Color(red, green, blue);
		g.setColor(randomColor);
		
		g.fillOval(this.getWidth() / 2, this.getHeight() / 2, 100, 100);
	}
}
