-- MySQL Script generated by MySQL Workbench
-- Mon Jan 18 15:29:30 2021
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema sissalud
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema sissalud
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `sissalud` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `sissalud` ;

-- -----------------------------------------------------
-- Table `sissalud`.`workdependency`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sissalud`.`workdependency` (
  `uuid` VARCHAR(36) NOT NULL,
  `name` VARCHAR(100) NULL DEFAULT NULL,
  PRIMARY KEY (`uuid`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `sissalud`.`job`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sissalud`.`job` (
  `uuid` VARCHAR(36) NOT NULL,
  `name` VARCHAR(100) NULL DEFAULT NULL,
  `description` VARCHAR(100) NULL DEFAULT NULL,
  PRIMARY KEY (`uuid`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `sissalud`.`person`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sissalud`.`person` (
  `uuid` VARCHAR(36) NOT NULL,
  `fullname` VARCHAR(100) NOT NULL,
  `cui` VARCHAR(70) NULL DEFAULT NULL,
  `job_uuid` VARCHAR(36) NULL DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  INDEX `person_job` (`job_uuid` ASC) VISIBLE,
  CONSTRAINT `person_job`
    FOREIGN KEY (`job_uuid`)
    REFERENCES `sissalud`.`job` (`uuid`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `sissalud`.`rol`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sissalud`.`rol` (
  `id` INT(11) NOT NULL,
  `role` VARCHAR(50) NULL DEFAULT NULL,
  `description` VARCHAR(50) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `sissalud`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sissalud`.`user` (
  `uuid` VARCHAR(36) NOT NULL,
  `username` VARCHAR(100) NOT NULL,
  `password` VARCHAR(200) NOT NULL,
  `rol_id` INT(11) NOT NULL,
  PRIMARY KEY (`uuid`),
  INDEX `user_rol` (`rol_id` ASC) VISIBLE,
  CONSTRAINT `user_rol`
    FOREIGN KEY (`rol_id`)
    REFERENCES `sissalud`.`rol` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `sissalud`.`autorization`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sissalud`.`autorization` (
  `uuid` VARCHAR(36) NOT NULL,
  `register` INT(11) NOT NULL,
  `dateemmited` DATE NULL DEFAULT NULL,
  `startdate` DATE NULL DEFAULT NULL,
  `enddate` DATE NULL DEFAULT NULL,
  `resumework` DATE NULL DEFAULT NULL,
  `totaldays` INT(11) NULL DEFAULT NULL,
  `pendingdays` INT(11) NULL DEFAULT NULL,
  `observation` TEXT NULL DEFAULT NULL,
  `authorizationyear` VARCHAR(4) NULL DEFAULT NULL,
  `person_idperson` VARCHAR(36) NOT NULL,
  `partida` VARCHAR(45) NULL,
  `workdependency_uuid` VARCHAR(36) NOT NULL,
  `user_uuid` VARCHAR(36) NOT NULL,
  PRIMARY KEY (`uuid`, `register`),
  INDEX `autorization_person` (`person_idperson` ASC) VISIBLE,
  INDEX `autorization_user` (`user_uuid` ASC) VISIBLE,
  INDEX `fk_autorization_workdependency1_idx` (`workdependency_uuid` ASC) VISIBLE,
  CONSTRAINT `autorization_person`
    FOREIGN KEY (`person_idperson`)
    REFERENCES `sissalud`.`person` (`uuid`),
  CONSTRAINT `autorization_user`
    FOREIGN KEY (`user_uuid`)
    REFERENCES `sissalud`.`user` (`uuid`),
  CONSTRAINT `fk_autorization_workdependency1`
    FOREIGN KEY (`workdependency_uuid`)
    REFERENCES `sissalud`.`workdependency` (`uuid`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
