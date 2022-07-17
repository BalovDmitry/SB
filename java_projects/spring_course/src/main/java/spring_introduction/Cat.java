package spring_introduction;

import org.springframework.stereotype.Component;

//@Component("catBean")
public class Cat implements Pet 
{
	public Cat()
	{
		System.out.println("Cat was created");
	}
	
	@Override
	public void say() 
	{
		System.out.println("Meow-meow");
	}

}
