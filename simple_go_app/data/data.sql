INSERT INTO todo (title, description)
SELECT * FROM (
    VALUES  ('First todo', 'This is the first todo'),
            ('Second todo', 'This is the second todo'),
            ('Third todo', 'This is the third todo'),
            ('Fourth todo', 'This is the fourth todo'),
            ('Fifth todo', 'This is the fifth todo')) src
WHERE NOT EXISTS (SELECT 1 FROM todo);