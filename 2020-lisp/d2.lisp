(defun read-to-list (path)
  (with-open-file (f path)
	(loop for line = (read-line f nil)
		  while line
		  collect line)))

(defstruct entry min max chr passwd)

(defun parse (line)
	(let* ((col (search ":" line))
		  (dash (search "-" line))
		  (space (search " " line))
		  (min (subseq line 0 dash))
		  (max (subseq line (+ 1 dash) space))
		  (chr (char line (+ 1 space)))
		  (passwd (subseq line (+ 2 col))))
			(make-entry 
			  :min (parse-integer min) 
			  :max (parse-integer max) 
			  :chr chr
			  :passwd passwd)))

(defun parse-data (data) (mapcar #'parse data))

(defun valid-p1 (e)
  (let ((cnt (count (entry-chr e) (entry-passwd e))))
	(and (<= (entry-min e) cnt)
		 (>= (entry-max e) cnt))))

(defun xor (a b)
	(or (and a (not b))
		(and b (not a))))

(defun valid-p2 (e)
	(let ((c1 (char (entry-passwd e) (- (entry-min e) 1)))
		  (c2 (char (entry-passwd e) (- (entry-max e) 1)))
		  (c (entry-chr e)))
		  (xor (eql c1 c) (eql c2 c))))

(defun solve-p1 (data) (count-if #'valid-p1 data))
(defun solve-p2 (data) (count-if #'valid-p2 data))

(let ((data (parse-data (read-to-list "d2.txt"))))
	(assert (= 500 (solve-p1 data)))
	(assert (= 313 (solve-p2 data))))

