/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//275 area codes
var usAreaCodes []int = []int{201, 202, 203, 205, 206, 207, 208, 209, 210, 212, 213, 214, 215, 216, 217, 218, 219, 224, 225, 228, 229, 231, 234, 239, 240, 248, 251, 252, 253, 254, 256, 260, 262, 267, 269, 270, 276, 281, 301, 302, 303, 304, 305, 307, 308, 309, 310, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 323, 325, 330, 331, 334, 336, 337, 339, 347, 351, 352, 360, 361, 386, 401, 402, 404, 405, 406, 407, 408, 409, 410, 412, 413, 414, 415, 417, 419, 423, 424, 425, 430, 432, 434, 435, 440, 443, 469, 478, 479, 480, 484, 501, 502, 503, 504, 505, 507, 508, 509, 510, 512, 513, 515, 516, 517, 518, 520, 530, 540, 541, 551, 559, 561, 562, 563, 567, 570, 571, 573, 574, 575, 580, 585, 586, 601, 602, 603, 605, 606, 607, 608, 609, 610, 612, 614, 615, 616, 617, 618, 619, 620, 623, 626, 630, 631, 636, 641, 646, 650, 651, 660, 661, 662, 678, 682, 701, 702, 703, 704, 706, 707, 708, 712, 713, 714, 715, 716, 717, 718, 719, 720, 724, 727, 731, 732, 734, 740, 754, 757, 760, 762, 763, 765, 769, 770, 772, 773, 774, 775, 781, 785, 786, 801, 802, 803, 804, 805, 806, 808, 810, 812, 813, 814, 815, 816, 817, 818, 828, 830, 831, 832, 840, 843, 845, 847, 848, 850, 856, 857, 858, 859, 860, 862, 863, 864, 865, 870, 878, 901, 903, 904, 906, 907, 908, 909, 910, 912, 913, 914, 915, 916, 917, 918, 919, 920, 925, 928, 931, 936, 937, 941, 947, 949, 951, 952, 954, 956, 970, 971, 972, 973, 978, 979, 980, 985, 989}
var Names []string = []string{"Brenda", "Henderson", "Julia", "Mitchell", "Lisa", "Romero", "Michael", "Cooper", "Samantha", "Webb", "Daryl", "Davis", "Christine", "Bender", "Toni", "Smith", "Martin", "Delacruz", "Timothy", "Payne", "Tamara", "Smith", "Cathy", "Huynh", "Tom", "Wolfe", "Gina", "Mason", "Sarah", "Torres", "Victoria", "Downs", "Jordan", "Wong", "Raymond", "Thornton", "Scott", "Phelps", "Kristin", "Holmes", "Crystal", "Allen", "Steven", "Browning", "Sarah", "Richardson", "William", "Jones", "Michael", "Garcia", "Nathan", "Rice", "Mrs.", "Michele", "Julie", "Hill", "Nicole", "Cochran", "Edward", "Wright", "Robert", "Lee", "Charles", "Johnson", "Michael", "Rodriguez", "David", "Larson", "Robert", "Mccall", "Denise", "Ray", "Joseph", "Mendoza", "Nicholas", "Miller", "Jared", "Wyatt", "Marisa", "Perez", "Sharon", "Young", "Luis", "Scott", "Jessica", "Duran", "Michael", "Williams", "Sean", "Roberts", "Kevin", "Cain", "Lauren", "Barnes", "Andrea", "Singh", "Russell", "Bradley", "Holly", "Watson", "Laura", "Johnson", "Dale", "Hicks", "Ashley", "Brown", "Shawn", "Sanders", "Anna", "Elliott", "Edward", "Moore", "Michael", "Butler", "Courtney", "Young", "Ricardo", "Waters", "Andrea", "Mcfarland", "Sara", "Cook", "Kevin", "Turner", "Matthew", "Gibson", "Bonnie", "Cummings", "Dennis", "Hawkins", "Phillip", "Smith", "Elizabeth", "Tanner", "Stacy", "Walters", "Katie", "Simpson",
	"Anthony", "Ayers", "Jennifer", "Snyder", "Richard", "Rose", "Kimberly", "Berry", "Julie", "Martinez", "Holly", "Logan", "Andrew", "Brown", "Jacob", "Larsen", "Eric", "Graves", "Catherine", "Richard", "David", "Horton", "Judith", "Obrien", "Amanda", "King", "Erika", "Moreno", "Charles", "Rodriguez", "Trevor", "Roth", "Dr.", "David", "Melanie", "Navarro", "Jose", "Castillo", "Joseph", "Davis", "Kimberly", "Mendez", "Anthony", "James", "Erin", "Reyes", "Alexander", "Patterson", "Mark", "Nguyen", "Joseph", "Green", "Marcus", "Sharp", "Jay", "Houston", "Megan", "Alvarado", "Brenda", "Walker", "Gina", "Mahoney", "Amy", "Cummings", "Teresa", "Phillips", "Alexander", "Jones", "Steven", "Taylor", "Katie", "Hayes", "Lorraine", "Fischer", "Rachel", "Mccullough", "Melissa", "Myers", "Emily", "Fowler", "Julie", "Le", "Abigail", "Stewart", "Greg", "Jensen", "Rachel", "Perkins", "Abigail", "Lyons", "Teresa", "Stewart", "Shelby", "Miller", "James", "Griffith", "Brenda", "Johnson", "Carla", "Yates", "George", "Williams", "Jason", "Beasley", "Kyle", "Davis", "Melissa", "Rhodes", "Carolyn", "Wiley", "Allen", "Green", "Tom", "Martin", "Annette", "Frazier", "David", "Wilson", "Kathy", "Hayes", "Jo", "Johns", "Elizabeth", "Baker", "Donald", "Perez", "Harold", "Espinoza", "John", "Baker", "Rhonda", "Smith", "Veronica", "Farrell", "Devin", "Russell", "Teresa", "Johnson", "Henry", "Johnson", "Kimberly",
	"Sutton", "John", "Zuniga", "Jessica", "Daniels", "Derek", "Baker", "Rachel", "Howell", "Michelle", "Wood", "Colleen", "Harris", "Anita", "Weaver", "Kristy", "Fernandez", "Richard", "Adams", "Ricardo", "Carter", "Courtney", "Martin", "Adrian", "Cooper", "James", "Ortiz", "Travis", "Mcclain", "Johnny", "Hoffman", "Emily", "Mcdaniel", "Kimberly", "Fitzgerald", "Jose", "Copeland", "Austin", "Holder", "Terry", "Maxwell", "Victor", "Hall", "David", "Lane", "Clarence", "Sherman", "William", "Bush", "Latoya", "Olsen", "Carmen", "Rice", "Nathan", "Barrett", "Shawn", "Martin", "Jamie", "Freeman", "Michael", "Young", "Andrew", "Burns", "Zachary", "Gonzalez", "Kimberly", "Newman", "Crystal", "Cole", "Angelica", "Schmidt", "Nancy", "Cole", "Andrew", "Flores", "Stephen", "Smith", "Kaitlyn", "Jones", "Timothy", "Rowe", "Philip", "Andrade", "Nicholas", "Mckenzie", "Laura", "Mendez", "Tonya", "Allison", "Janice", "Cole", "Rebecca", "Myers", "Alex", "Park", "Jennifer", "Lee", "Sarah", "Hall", "Margaret", "Vasquez", "Ashley", "Foley", "Shawna", "Osborne", "Jeffrey", "Jenkins", "Colleen", "Salazar", "William", "Atkinson", "James", "Burns", "Daniel", "Greene", "Tracy", "Rocha", "Diana", "Underwood", "Kurt", "Chan", "Robert", "Cowan", "David", "Atkinson", "Francis", "Martin", "Keith", "English", "Michael", "Andrews", "Kari", "Atkinson", "Emily", "Bowman", "Kenneth", "Miller", "Luke", "Goodman", "April",
	"Roberts", "Lisa", "Brown", "Gary", "Lopez", "Edward", "Edwards", "Bethany", "Ramirez", "Anthony", "Jones", "Elizabeth", "Garcia", "David", "Ramirez", "Meghan", "Parker", "Julie", "Williams", "John", "Murphy", "Kristin", "Marshall", "Cory", "Patterson", "Roy", "Martin", "Brittany", "Hutchinson", "Matthew", "Ray", "Christopher", "Owens", "Tammy", "Williams", "Christian", "Peterson", "Kevin", "Gray", "Dennis", "Williams", "Linda", "Harrison", "Cassandra", "Hines", "Matthew", "Hopkins", "Joyce", "Orr", "James", "Silva", "Fernando", "Jones", "Chelsea", "Simpson", "Stacy", "Cuevas", "Benjamin", "Schwartz", "Jennifer", "Carpenter", "Heather", "Jones", "Margaret", "Franklin", "Stephanie", "Thompson", "Leah", "Herrera", "Mark", "Mendoza", "Nathaniel", "Richardson", "Patrick", "Walter", "Corey", "Sanders", "Amy", "Pierce", "Thomas", "Holt", "Jacqueline", "Todd", "Aaron", "Spencer", "Katherine", "Coleman", "Susan", "Berger", "Nicholas", "Chen", "Michael", "Anderson", "Julia", "Wright", "Sarah", "Ashley", "Christopher", "Sanchez", "Dr.", "Ricky", "Christopher", "Hernandez", "William", "Griffith", "Derek", "Hardin", "Mr.", "Todd", "Theresa", "Rivera", "Aaron", "Yang", "Angela", "Bradley", "Barbara", "Reyes", "Chelsea", "Jones", "Amanda", "Martinez", "Shawna", "Murphy", "Holly", "Lewis", "Danny", "Blanchard", "George", "Travis", "Abigail", "Meyer", "Paige", "Beard", "Ronnie", "Landry", "James",
	"Martin", "John", "Thomas", "Susan", "Williams", "Michael", "Miles", "Christopher", "Pratt", "Sarah", "Kelly", "John", "Ramos", "Andrea", "Hooper", "Frank", "Johnson", "Joshua", "Walton", "Lisa", "White", "Ashley", "Garcia", "Kari", "White", "Veronica", "Martin", "Michael", "White", "Brittney", "Grant", "Matthew", "Hernandez", "Tyler", "Jones", "Grant", "Figueroa", "Sierra", "Bailey", "Samuel", "Torres", "Catherine", "Allen", "Amy", "Morales", "Ryan", "Rojas", "Justin", "Wall", "Michelle", "Smith", "Sarah", "Perez", "Maurice", "Berger", "Phillip", "Wood", "Justin", "Martin", "Gregory", "Lloyd", "David", "Osborne", "Meghan", "Simpson", "Kathryn", "Valdez", "Casey", "Wallace", "Eric", "Allison", "Robert", "Stewart", "Stephanie", "Palmer", "Matthew", "Stone", "Kari", "Greene", "Mark", "Figueroa", "Bridget", "Buckley", "Brian", "Martin", "Cynthia", "Gaines", "Christopher", "Miller", "Jeffrey", "Avery", "Danielle", "Vincent", "Robert", "Jackson", "Marissa", "George", "Joshua", "Hobbs", "Jessica", "Diaz", "David", "Gutierrez", "Timothy", "Taylor", "Deanna", "Griffin", "Jacqueline", "Norton", "Cheryl", "Sanchez", "Ashley", "Ibarra", "Jennifer", "Ellis", "Lauren", "Mcdonald", "Richard", "Wells", "Mckenzie", "Hernandez", "Victoria", "Rasmussen", "Nicholas", "Allen", "Dawn", "Ramirez", "Deanna", "Bailey", "Diane", "Curry", "Michael", "Bean", "Joe", "Glenn", "William", "Flores", "Danielle",
	"Jimenez", "Patricia", "Anderson", "Angela", "Lopez", "Janice", "Allen", "Erin", "Riley", "Krista", "Burke", "Charles", "Pena", "Melissa", "Barton", "Allison", "Lee", "Laura", "Phillips", "Kelly", "Anderson", "Maria", "Watson", "Melissa", "Anderson", "Jessica", "Barnes", "Rachel", "Wilson", "Gabriella", "Wilson", "Sean", "Cole", "Timothy", "Hernandez", "Carol", "Ramos", "Anthony", "Fisher", "Christina", "Hall", "Kristi", "Carpenter", "Matthew", "Boyer", "Randy", "Lucas", "Donna", "Robinson", "Laura", "Alvarado", "Courtney", "Smith", "Scott", "Fields", "Pamela", "Hall", "Lauren", "Smith", "Nicholas", "Rivera", "Christopher", "Bell", "Christine", "Silva", "Teresa", "Shepherd", "Mr.", "Aaron", "Sandra", "Lopez", "Robert", "Wu", "Chris", "Friedman", "Michael", "Flores", "Tommy", "Sandoval", "Ann", "Rivera", "Mary", "Woods", "Jennifer", "Turner", "Joseph", "Harvey", "Tiffany", "Sampson", "James", "Oconnell", "James", "Mejia", "Carolyn", "White", "Brandi", "Wallace", "Jordan", "Howe", "Jenna", "Thomas", "Kiara", "Moody", "Shaun", "Martinez", "Jose", "Higgins", "Amy", "Allen", "Martha", "Miller", "Katherine", "Weeks", "Christine", "Yates", "Luis", "Douglas", "Nicole", "Price", "Kevin", "Howell", "Alicia", "Scott", "Hayley", "Hernandez", "Sally", "Powers", "Nancy", "Salazar", "Mark", "Powers", "Lindsay", "Ayers", "Steven", "Hess", "Elizabeth", "Merritt", "Matthew", "Stephens", "Nathan",
	"Wilkinson", "Mrs.", "Courtney", "Douglas", "Long", "Mary", "Morton", "Dana", "Bartlett", "Lisa", "Torres", "Krystal", "Johnston", "Bradley", "Barker", "Brenda", "Hunt", "Joseph", "Smith", "Patricia", "Smith", "William", "Duarte", "James", "Cordova", "Brian", "Parks", "Melissa", "Garcia", "Amber", "Mayo", "Katherine", "Jones", "Alexander", "Blair", "Christopher", "Watson", "Alison", "James", "Norman", "Smith", "Jacob", "Reed", "Timothy", "Larson", "Dana", "Morton", "Cheyenne", "Green", "Jonathan", "Hamilton", "Heather", "Holmes", "Jason", "Perez", "Kevin", "Roberts", "Michael", "Rowland", "Emma", "Young", "Nicole", "Garcia", "Danielle", "Green", "Margaret", "Robinson", "Alan", "Hawkins", "William", "Green", "Laura", "Gomez", "Justin", "Rasmussen", "Mrs.", "Kimberly", "Heather", "Taylor", "Wesley", "Kennedy", "Daniel", "Cook", "Fernando", "Campbell", "Deborah", "Foley", "Victor", "Lewis", "Taylor", "Newman", "Andrew", "Hicks", "Gregory", "Osborne", "Shari", "Beasley", "Sarah", "Smith", "Christian", "Jackson", "Chelsea", "Barnes", "Jeremy", "Castillo", "Michelle", "Brandt", "Sherry", "Mcdonald", "Jennifer", "Douglas", "Tyler", "Gutierrez", "Michelle", "Murillo", "Michelle", "Dixon", "Lawrence", "Hogan", "Blake", "Travis", "Jared", "Ortiz", "Manuel", "Barnes", "Tammy", "Ramirez", "Stephen", "Lewis", "Michael", "Combs", "Beth", "Johnson", "Mariah", "Blankenship", "Gregg", "West",
	"Andrew", "Montgomery", "Brandon", "Benitez", "Randall", "Lewis", "Samantha", "Francis", "Dana", "Thomas", "Ashley", "Oneal", "Fernando", "Brown", "Stephanie", "Stevenson", "Joseph", "Green", "Christian", "Chambers", "Renee", "Brady", "Darren", "Sweeney", "Timothy", "Hodges", "Steve", "Mathews", "Christine", "Chandler",
	"Michael", "Ruiz", "William", "Martin", "Joshua", "Kramer", "Michael", "Roman", "Ronnie", "Andrews", "Douglas", "Young", "Robert", "Smith", "Michael", "Jackson", "Robert", "Wagner", "Cory", "Young", "Kevin", "Nguyen", "Jon", "Lara", "Donald", "Flores", "Jason", "Cox", "Mark", "Allen", "Joe", "Smith", "Eric", "Austin", "Edwin", "Davis", "Zachary", "Combs", "Michael", "Rhodes", "Jason", "Smith", "Mark", "Carr", "Timothy", "Smith", "Lee", "Taylor", "Bruce", "Gibson", "Matthew", "Villanueva", "John", "Mcintyre", "William", "Henry", "Eric", "Murphy", "Barry", "Hernandez", "Alexander", "Bridges", "Chad", "Scott", "Leslie", "Mcdonald", "Christopher", "Campbell", "Brian", "Davidson", "Samuel", "Mooney", "Andrew", "Morgan", "David", "Thomas", "Jimmy", "Williams", "Bradley", "Stone", "Nathan", "Newman", "Victor", "Roberts", "Anthony", "Keith", "William", "Camacho", "Cory", "Yu", "Bradley", "Hughes", "Johnathan", "Weaver", "Jason", "Miles", "Dylan", "Rowe", "James", "Ramirez", "Jeremiah", "Gomez", "Darrell", "Harvey", "John", "Gonzalez", "John", "Edwards", "Sean", "Scott", "Charles", "Nelson", "Mark", "Johnson", "William", "Wright", "Bryan", "Tran", "Anthony", "Mccormick", "Gregory", "Gonzales", "Dustin", "Molina", "Nathan", "Moss", "Bobby", "Anderson", "Lance", "Colon", "Edward", "Cisneros", "Jason", "Perry", "Gerald", "Garcia", "Paul", "Chavez", "Adam", "Walker", "Michael", "Garcia",
	"Tony", "Jones", "Brandon", "Turner", "Jeffrey", "Lopez", "William", "Velez", "Jeremy", "Bowman", "Christopher", "Smith", "Harold", "Thompson", "Taylor", "Mcmahon", "Timothy", "Davila", "Jason", "Odom", "Richard", "Cardenas", "William", "Jones", "Jerome", "Hudson", "Brian", "Harris", "Tony", "Sullivan", "Timothy", "Johnson", "Scott", "Andersen", "Brian", "Velazquez", "Jeffery", "Padilla", "Michael", "Collins", "Adam", "Sanchez", "Timothy", "Rodriguez", "Michael", "Patrick", "Donald", "Williamson", "Kyle", "Bailey", "Charles", "Alexander", "Justin", "Hernandez", "Phillip", "Mcgee", "Jason", "Price", "Brian", "Welch", "Jason", "Khan", "Brett", "Davis", "William", "Williams", "Antonio", "Clay", "Robert", "Davis", "Benjamin", "Young", "Erik", "Lynch", "Jamie", "Kennedy", "Ronald", "Garrison", "Joshua", "Saunders", "Kevin", "Thornton", "Eric", "Clark", "Benjamin", "Dillon", "Christopher", "Alexander", "Michael", "Silva", "Tyrone", "Castaneda", "Richard", "Braun", "Joe", "Wood", "Ian", "Arellano", "David", "Powell", "Steven", "Owens", "David", "Stanley", "Jeffrey", "Delgado", "George", "Palmer", "Allen", "Hill", "Ryan", "Guzman", "Daniel", "Scott", "Johnathan", "Macias", "Lee", "Barry", "Alan", "Melendez", "David", "Schwartz", "Christopher", "Trevino", "Brian", "Stephens", "Travis", "Anderson", "Eric", "Banks", "William", "Lucas", "David", "Oconnor", "Christopher", "Newton", "Andrew",
	"Ferguson", "Daniel", "Johnson", "Anthony", "Miller", "David", "White", "Stephen", "Hahn", "Steven", "Barajas", "Matthew", "Mays", "Samuel", "May", "Daniel", "Flores", "John", "Watson", "Miguel", "Anderson", "Chad", "Nichols", "Peter", "Adams", "Scott", "Cuevas", "Dwayne", "Cabrera", "David", "Dixon", "Joshua", "Meyers", "Mark", "Wall", "Robert", "Warren", "Ross", "Brown", "Matthew", "Aguilar", "Joseph", "Bennett", "Nicholas", "Gray", "Evan", "Spencer", "David", "Campbell", "David", "Thompson", "Chad", "Peterson", "Kevin", "Jones", "David", "Phillips", "Thomas", "Mason", "Warren", "Tran", "Aaron", "Beard", "Kurt", "Garcia", "Roberto", "Hamilton", "Gary", "Oconnor", "Paul", "Knight", "Lee", "Atkinson", "Bryan", "Brown", "Jason", "Johnson", "Philip", "Ortiz", "Lance", "Schroeder", "Dr.", "Justin", "Joseph", "Rose", "Andre", "Jackson", "Craig", "Marshall", "Robert", "Brown", "Ruben", "Fry", "Matthew", "Perkins", "Kyle", "Carson", "Alejandro", "Thomas", "Mark", "Olson", "Carlos", "Hanson", "Gregory", "Villegas", "Joshua", "Wu", "Roy", "Garcia", "Christopher", "Torres", "George", "Koch", "Jose", "Garner", "Robert", "Horton", "Daniel", "Nichols", "Charles", "Leon", "Joshua", "Stevens", "Kevin", "Hicks", "John", "Porter", "Zachary", "Hess", "Paul", "Mendez", "Kevin", "Soto", "Matthew", "Castillo", "Anthony", "Gonzales", "Justin", "Rivas", "Mark", "Taylor", "Ronald", "Davis", "John",
	"White", "Dylan", "Johnson", "Larry", "Ford", "Raymond", "Ryan", "Michael", "Moreno", "Lawrence", "Pham", "Wayne", "Pierce", "Michael", "Arias", "Stephen", "Smith", "Matthew", "Osborne", "Brandon", "Brown", "Anthony", "Graham", "Adam", "Shannon", "James", "Rodriguez", "Gabriel", "Li", "Danny", "Morris", "Scott", "Garcia", "Nicholas", "Combs", "Ryan", "Thomas", "Benjamin", "Moore", "Michael", "Johnson", "Trevor", "Waller", "Manuel", "Mckinney", "Brady", "Roberts", "Joseph", "Haas", "Jonathan", "Gutierrez", "Mark", "Turner", "Todd", "Murray", "Brandon", "Johnson", "Ronald", "Long", "John", "Wong", "Justin", "Hopkins", "Samuel", "Beck", "Matthew", "Moon", "Andrew", "Williams", "Bradley", "Parsons", "Brent", "Clarke", "Kevin", "Allen", "Michael", "Clark", "Jeremy", "Martin", "Bill", "Mann", "James", "Tanner", "Robert", "Burgess", "Charles", "Sanchez", "Jason", "Ramos", "Michael", "Williams", "Kyle", "Riley", "James", "Nelson", "Justin", "Lee", "Brent", "Chavez", "Joseph", "Bolton", "Phillip", "Abbott", "Kevin", "Craig", "James", "Davis", "Roger", "Bradley", "Nathan", "Johnson", "Matthew", "Bean", "Henry", "Combs", "Harry", "Baxter", "Thomas", "Jackson", "Wesley", "Johnson", "David", "Gonzales", "Dylan", "Love", "Nicholas", "Miller", "Rodney", "Owens", "Nathaniel", "Williams", "Roger", "Norman", "Jonathon", "Malone", "Eric", "Taylor", "Cameron", "Chang", "Travis", "Scott", "Marco",
	"Wilson", "Gregory", "Ray", "John", "Montes", "David", "Walters", "Robert", "Huang", "Brian", "Smith", "Scott", "Schneider", "Lawrence", "Howard", "Marcus", "Mooney", "Alfred", "Scott", "Ryan", "Bonilla", "Joseph", "Bowen", "Michael", "Nelson", "Erik", "Mason", "Michael", "Diaz", "Kevin", "Terry", "Jay", "Aguilar", "James", "Montgomery", "Isaac", "Allison", "Luis", "Thomas", "John", "Rios", "James", "Andrews", "Roberto", "Hicks", "Paul", "Rodriguez", "Robert", "Williams", "John", "Snyder", "John", "Greene", "Jacob", "Warren", "Richard", "Salas", "Bobby", "Jones", "Christopher", "Lee", "Karl", "Cuevas", "Julian", "Copeland", "Derek", "Allen", "Joseph", "Rodriguez", "Ronald", "Torres", "Max", "Spencer", "Jacob", "Frazier", "Oscar", "Hunter", "Steven", "Washington", "Patrick", "Espinoza", "Jared", "Gonzales", "Eric", "Christensen", "Jorge", "Wilson", "Zachary", "Powell", "Justin", "Daniel", "Christopher", "Hudson", "Joshua", "Michael", "Christopher", "Greene", "Matthew", "Hernandez", "James", "Zuniga", "Luis", "Carter", "Luke", "Robles", "Larry", "Diaz", "Jorge", "Adkins", "Ian", "Turner", "Jeremy", "Flores", "James", "Rivera", "Scott", "Brown", "John", "Carter", "Timothy", "Morris", "Ronald", "Price", "Mr.", "James", "Peter", "Nash", "Alejandro", "Hammond", "Brian", "Newman", "Kyle", "White", "Steven", "Mcdonald", "Jonathan", "Downs", "Matthew", "Franklin", "Joshua", "Henry",
	"Nathaniel", "Ellis", "Brandon", "Mata", "Carl", "Berg", "Robert", "Reed", "Andrew", "Rice", "Scott", "Wiley", "Edward", "Thompson", "Luis", "Hall", "Paul", "Russell", "Dustin", "Short", "Tyler", "Allen", "Kevin", "Wall", "Cory", "Murillo", "Michael", "Hampton", "Ryan", "Gallegos", "Alexander", "Johnson", "Brian", "Hines", "Ray", "Duran", "Carl", "Swanson", "William", "Williams", "Kenneth", "Chandler", "Alexander", "Massey", "John", "Abbott", "Charles", "Pena", "James", "Jordan", "Scott", "Jordan", "Francis", "Jimenez", "Kyle", "Russell", "Roger", "Pierce", "Franklin", "Miller", "Eric", "Harmon", "Phillip", "Cooper", "Jeffery", "Nguyen", "Tyler", "Cummings", "Tyler", "Torres", "Anthony", "Jackson", "Joseph", "Smith", "Mitchell", "Porter", "Donald", "Wilson", "Joseph", "Galvan", "Scott", "Middleton", "Jack", "Landry", "Steve", "Noble", "Gary", "Bowen", "James", "Barnes", "Mr.", "Timothy", "John", "Yang", "Edward", "Rangel", "Jeremy", "Harper", "Matthew", "Knight", "Andrew", "Tapia", "Peter", "Page", "Andrew", "Bailey", "Michael", "Johnson", "Jeffrey", "Hill", "William", "Walker", "Tommy", "Ellison", "Michael", "Davis", "Jose", "Reed", "Jonathon", "Peterson", "Clinton", "Thomas", "Stephen", "White", "Phillip", "Carlson", "Sean", "Warren", "Patrick", "Turner", "Patrick", "Thompson", "Mark", "Alexander", "Robert", "Vaughn", "Isaac", "Rivera", "Levi", "Bailey", "Caleb", "Nguyen",
	"Stephen", "Wallace", "Michael", "Wade", "Corey", "Gallegos", "Michael", "Pearson", "Nicholas", "Lopez", "Christopher", "Escobar", "Mr.", "Aaron", "Evan", "Taylor", "Brandon", "Lopez", "Steven", "Neal", "Tom", "Erickson", "John", "Ramirez", "Tommy", "Murphy", "Gary", "Baker", "Peter", "Pena", "Robert", "Allen", "Eric", "Hill", "Adam", "Jones", "Xavier", "Singleton", "Jerome", "Chen", "Benjamin", "Ward", "Joshua", "Edwards", "Jeffery", "Sharp", "Ryan", "Wood", "Kenneth", "Bryant", "William", "Tapia", "Jeff", "Smith", "Scott", "Carpenter", "Tyler", "Morris", "Keith", "Joyce", "David", "Mcdaniel", "Matthew", "Strickland", "Kevin", "Reed", "Daniel", "Hill", "Nicholas", "Gross", "Antonio", "Graves", "Joseph", "Martin", "Michael", "Brady", "Keith", "Brown", "Nicholas", "Peters", "Justin", "Roberts", "Christopher", "Gutierrez", "Brian", "Reed", "Brandon", "Martin", "Jacob", "May", "Jonathan", "Franco", "Keith", "Lynch", "Connor", "Moore", "Michael", "Atkinson", "Kenneth", "Graham", "Stephen", "Gilbert", "Jose", "Montes", "Adam", "Raymond", "Andrew", "Chavez", "Tyler", "Chavez", "Gary", "Rodriguez", "Jason", "Shaffer", "Luis", "Brady", "Jason", "Rosales", "Jeffrey", "Cunningham", "Hector", "Hernandez", "Christopher", "Vega", "Bill", "Nelson", "Joseph", "Anderson", "Jeffery", "West", "Christopher", "Morrison", "James", "Johnson", "David", "Peters", "Michael", "Barnes", "Samuel",
	"Brandt", "Eric", "Martin", "John", "Jackson", "Dean", "Hood", "Carlos", "Kirk", "Jimmy", "Lucero", "Nathaniel", "Walton"}

var uniqueNums map[string]bool = make(map[string]bool)

func getRandUsAreaCodes() int {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	return usAreaCodes[r.Intn(275)]
}

func getRandNum(unique bool) string {
	//5550001-5559999
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := fmt.Sprintf("%d%d", getRandUsAreaCodes(), r.Intn(9998)+5550001)
	if unique {
		if _, ok := uniqueNums[n]; !ok {
			uniqueNums[n] = true
			return n
		} else {

			for {
				n := fmt.Sprintf("%d%d", getRandUsAreaCodes(), r.Intn(9998)+5550001)
				if _, ok := uniqueNums[n]; !ok {
					uniqueNums[n] = true
					return n
				}
			}
		}
	}
	return n
}

func getRandAge() string {
	//5550100, 5550199+1
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	return fmt.Sprintf("%d", r.Intn(99))
}

func getRandName() string {
	//5550100, 5550199+1
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	return Names[r.Intn(len(Names))]
}

func GenerateRandNums(limit uint64, unique bool) []string {
	println("generating random Phone Numbers")
	var nums []string
	for i := uint64(0); i < limit; i++ {
		nums = append(nums, getRandNum(unique))
	}
	return nums
}

func GenerateRandNames(limit uint64) []string {
	println("generating random Names")
	var names []string
	for i := uint64(0); i < limit; i++ {
		names = append(names, getRandName())
	}
	return names
}

func GenerateRandAges(limit uint64) []string {
	println("generating random ages")
	var ages []string
	for i := uint64(0); i < limit; i++ {
		ages = append(ages, getRandAge())
	}
	return ages
}

func WriteCsv(name string, limit uint64, data ...[]string) {
	csvFile, err := os.Create(fmt.Sprintf("%s.csv", name))
	csvwriter := csv.NewWriter(csvFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer csvFile.Close()

	output := make([][]string, limit)
	for i := range output {
		output[i] = make([]string, len(data))
	}

	for i := 0; i < len(data); i++ {
		for j, d := range data[i] {
			output[j][i] = d
		}
	}
	csvwriter.WriteAll(output)

}

func GenerateData(cols []string, limit uint64, unique bool) [][]string {
	data := make([][]string, len(cols))

	for i, s := range cols {
		switch s {
		case "phone":
			data[i] = make([]string, limit)
			copy(data[i], GenerateRandNums(limit, unique))
		case "name":
			data[i] = make([]string, limit)
			copy(data[i], GenerateRandNames(limit))
		case "age":
			data[i] = make([]string, limit)
			copy(data[i], GenerateRandAges(limit))
		}
	}

	return data
}
