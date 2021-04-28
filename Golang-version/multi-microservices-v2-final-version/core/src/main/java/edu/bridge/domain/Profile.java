package edu.bridge.domain;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.Date;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Profile {
    private String username;    // 现名
    private Integer age;        // 年龄
    private Integer gender;     // 性别
    private String politicType; // 群众、共青团员、中共党员、民主人士等。
    private Date birthDate;     // 出生年月日
    private String photo;       // 照片
    private String oldName;     // 曾用名
    private String identityId;  // 身份证号
    private String career;      // 职业
    private String originHometown;// 原籍
    private String religion;    // 宗教信仰
    private String nation;      // 民族
    private String birthPlace;  // 出生地
    private String maritalStatus;// 婚姻状况
    private String homeAddress; // 家庭住址
    private String currentProblem;// 现实问题
    private String educationHistory;// 教育史
    private String healthState; // 健康状况
    private String maritalStatusExplicit;// 详细婚姻状况
    private String employmentHistory;// 就业史
    private String contractInstitutionHistory;// 以前与社会机构的接触
    private String hobby;       // 兴趣爱好
    private String otherSituation;// 其他需要说明的情况
}
