package spring_introduction;

import org.springframework.context.support.ClassPathXmlApplicationContext;

public class ScopeTest {

	public static void main(String[] args) {
		ClassPathXmlApplicationContext context = 
				new ClassPathXmlApplicationContext("applicationContext3.xml");
		
		Dog myDog = context.getBean("dogBean", Dog.class);
		myDog.say();
//		Dog yourDog = context.getBean("dogBean", Dog.class);
//		//myDog == yourDog ? System.out.println("Same object") : System.out.println("Different objects");
//		if (myDog == yourDog)
//		{
//			System.out.println("Same object");
//		}
//		else
//		{
//			System.out.println("Different objects");
//		}
		
		context.close();
	}

}
