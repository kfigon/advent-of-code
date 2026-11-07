(defun read-file-to-list (path)
  (with-open-file (f path)
		(loop for line = (read-line f nil)
			  while line
			  collect line)))

(defun parse (data) (mapcar #'parse-integer data))

(defun build-set (data)
	(let ((out (make-hash-table)))
	  (loop for d in data
			do (setf (gethash d out) t))
	  out))

(defun solve-p1 (data)
	(let ((s (build-set data)))
		 (loop for d in data
			do  (let ((to-find (- 2020 d)))
					(when (gethash to-find s)
						(return-from solve-p1 (* d to-find)))))))

(defun solve-p2 (data)
	(loop for i in data
		do (loop for j in data
				do (loop for k in data do 
										(when (= 2020 (+ i j k))
											  (return-from solve-p2 (* i j k)))))))

(assert (= 719796 (solve-p1 (parse (read-file-to-list "d1.txt")))))
(assert (= 144554112 (solve-p2 (parse (read-file-to-list "d1.txt")))))

