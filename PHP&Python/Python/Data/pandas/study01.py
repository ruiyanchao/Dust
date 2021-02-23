# -*- coding: UTF-8 -*-
import pandas as pd
import os


def main():
    path = os.getcwd()
    path = path+"/../../../../Tmp"
    df = pd.DataFrame({'工资': [100, 200, 300, 400], '绩效考核': [60, 70, 80, 90], '备注': ['不及格', '良好', '最佳', '优秀']}, index=['赵', '钱', '孙', '李'])
    print(df)
    print("----------打印 data frame-----------")
    df.to_csv(path+"/demo1.csv")
    df.to_excel(path+"/demo1.xlsx")
    print("------生成csv,xlsx--------")
    df2 = pd.read_csv(path+"/demo1.csv")
    df2.head(2)
    df3 = pd.read_excel(path+"/demo1.xlsx")
    print(df2)
    print(df3)


if __name__ == '__main__':
    main()
