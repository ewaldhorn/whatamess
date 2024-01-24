(ns leap)

(defn leap-year? [year] ;; <- argslist goes here
  ;; check that year % 4 == 0
  ;; if it is, it must not be % 100 == 0 but it must be % 400 == 0
  (and (zero? (mod year 4)) (or (zero? (mod year 400)) (pos? (mod year 100))))
)
