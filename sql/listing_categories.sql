SELECT c.catID, c.title, c.notes, count(m.qID)
  FROM Categories c
    LEFT JOIN CategoryMembership m
    ON c.catID = m.catID
  ORDER BY count(m.qID)
  ;
