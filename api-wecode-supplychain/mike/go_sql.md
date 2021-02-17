1. SELECT
ROW_NUMBER() OVER(ORDER BY document_number )
AS SEQUENECE,document_number,company_name AS SUPPLIER,seller_company_id,buyer_company_id,create_date ,expire_date ,delivery_date ,grand_total AS TOTAL,currency
FROM malar.dbo.document_header,malar.dbo.company
WHERE document_type = 120 AND buyer_company_id = id
AND company_type = 1003  AND currency = 'THB' AND document_status IN(1,3,5,7,9,11,13,15,17)AND 
create_date BETWEEN '2010-04-16 00:00:00' AND '2010-06-17 00:00:00' AND
expire_date BETWEEN '2010-04-16 00:00:00' AND '2010-06-17 00:00:00' AND
delivery_date BETWEEN '2010-04-16 00:00:00' AND '2010-06-17 00:00:00'


2. SELECT COUNT(document_number ) AS TOTAL
FROM malar.dbo.document_header,malar.dbo.company
WHERE document_type = 120 AND buyer_company_id = id
AND company_type = 1003 AND currency = 'THB' AND document_status IN(1,3,5,7,9,11,13,15,17)AND
create_date BETWEEN '2010-04-16 00:00:00' AND '2010-06-17 00:00:00' AND
expire_date BETWEEN '2010-04-16 00:00:00' AND '2010-06-17 00:00:00' AND
delivery_date BETWEEN '2010-04-16 00:00:00' AND '2010-06-17 00:00:00'