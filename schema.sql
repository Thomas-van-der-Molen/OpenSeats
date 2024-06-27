


drop table if exists reservations;
create table reservations(
	FName TINYTEXT,
	LName TINYTEXT,
	PNumber TINYTEXT,
	date TINYTEXT,
	RType TINYTEXT,
	quantity TINYTEXT,
	price TINYTEXT,
	internalNotes TINYTEXT
);

insert into reservations values 
("John", "Doe", "6014329876","2024-06-19", "Lounge Chair", "1", "45.00","check # 4193100"),
("James", "Smith", "1234567890", "2024-07-04", "High Top", "5", "75.00","check # 4193101"),
("Mary", "Aldrin", "1038396501", "2024-12-25" , "Day Bed", "1", "125.00","check # 4193102"),
("Mark", "Armstrong", "8148388375", "2023-01-01", "Main Deck Cabana", "1", "500.50","check # 4193103"),
("Mike", "Glenn", "9912168743", "2025-05-09", "Ocean View Cabana", "1", "750.99","check # 4193104");
