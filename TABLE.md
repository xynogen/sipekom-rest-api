# Entity Available on Database

## Table User

| Field       | Serialize Name | DataType |
| ----------- | -------------- | -------- |
| Username    | username       | string   |
| Password    | password       | string   |
| Photo       | photo          | string   |
| Role        | role           | uint8    |
| IsActivated | is_activated   | uint8    |

## Table PPDS

| Field       | Serialize Name | DataType |
| ----------- | -------------- | -------- |
| Name        | name           | string   |
| BirthDate   | birthdate      | string   |
| BirthPlace  | birthplace     | string   |
| NIK         | nik            | string   |
| NIM         | nim            | string   |
| Address     | address        | string   |
| PhoneNumber | phone_number   | string   |
| Angkatan    | angkatan       | uint     |
| Prodi       | prodi          | string   |
| STR         | str            | string   |
| SIP         | sip            | string   |
| Kompetensi  | kompetensi     | uint8    |
| IDUser      | id_user        | uint     |

## Table Konsulen

| Field     | Serialize Name | DataType |
| --------- | -------------- | -------- |
| Name      | name           | string   |
| Spesialis | spesialis      | string   |
| IDUser    | id_user        | uint     |

## Table Lokasi

| Field  | Serialize Name | DataType |
| ------ | -------------- | -------- |
| Lokasi | lokasi         | string   |
| Uri    | uri            | string   |

## Table ELogBook

| Field         | Serialize Name | DataType  |
| ------------- | -------------- | --------- |
| Title         | title          | string    |
| Jumlah        | jumlah         | uint      |
| StartTime     | start_time     | time.Time |
| EndTime       | end_time       | time.Time |
| Deskripsi     | deskripsi      | string    |
| MedicalRecord | medical_record | string    |
| IDUser        | id_user        | uint      |
| IDKonsulen    | id_konsulen    | uint      |

## Table Absensi

| Field     | Serialize Name | DataType |
| --------- | -------------- | -------- |
| Absen     | absen          | string   |
| AbsenFlag | absen_flag     | uint8    |
| Lokasi    | lokasi         | string   |
| IDUser    | id_user        | uint     |
