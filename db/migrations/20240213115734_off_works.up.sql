CREATE TABLE off_works (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    status ENUM('pending', 'accept', 'deny') DEFAULT 'pending',
    PRIMARY KEY (id)
)Engine=InnoDB