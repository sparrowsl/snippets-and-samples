CREATE TABLE IF NOT EXISTS employees(
  first_name  TEXT NOT NULL,
  last_name   TEXT NOT NULL,
  position    TEXT NOT NULL,
  date        DATE
);

INSERT INTO employees VALUES
  ('John', 'Johnson', 'Manager', '2016-12-31'),
  ('Tou', 'Xiong', 'Software Engineer', '2016-10-05'),
  ('Michaela', 'Michaelson', 'District Manager', '2015-12-19'),
  ('Jake', 'Jacobson', 'Programmer', null),
  ('Jacquelyn', 'Jackson', 'DBA', null),
  ('Sally', 'Weber', 'Web Developer', '2015-12-18');

-- First Name    Last Name     Position            Separation date
-- John          Johnson       Manager             2016-12-31
-- Tou           Xiong         Software Engineer   2016-10-05
-- Michaela      Michaelson    District Manager    2015-12-19
-- Jake          Jacobson      Programmer
-- Jacquelyn     Jackson       DBA
-- Sally         Weber         Web Developer       2015-12-18
