package spring_introduction;

import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Test4 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		ClassPathXmlApplicationContext context = 
				new ClassPathXmlApplicationContext("applicationContext2.xml");
		Dog myDog  = context.getBean("myPet", Dog.class);
		Dog yourDog  = context.getBean("myPet", Dog.class);
	
		System.out.println("Same oject references? " + (myDog == yourDog));
		System.out.println(myDog);
		System.out.println(yourDog);
		
//		myDog.setName("Belka");
//		yourDog.setName("Strelka");
//		System.out.println(myDog.getName());
//		System.out.println(yourDog.getName());
		context.close();
	}

}
