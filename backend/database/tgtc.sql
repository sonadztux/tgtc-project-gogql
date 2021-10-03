CREATE TABLE IF NOT EXISTS "User" ("ID" int PRIMARY KEY, Name text, tier text);

CREATE TABLE IF NOT EXISTS "Coupon" ("ID" int PRIMARY KEY, Name text, Type text, Status text, Amount float, image_url text, usertier text, start_date date, expire_date date);

CREATE TABLE IF NOT EXISTS "DetailCoupon" (UserID int, CouponID int, FOREIGN KEY(UserID) REFERENCES "User"("ID"), FOREIGN KEY(CouponID) REFERENCES "Coupon"("ID"));