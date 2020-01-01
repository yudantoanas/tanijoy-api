
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO project_files (project_id, filename, path, label, purpose) VALUES
(7, 'Prospektus Cabai Rawit Bogor', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/documents/Dokumen-Investasi-Cabai-Rawit-Bogor.pdf', 'document', 'prospektus'),
(8, 'Prospektus Kentang Granola Berastagi', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/documents/Dokumen-Investasi-Kentang-Granola-Berastagi.pdf', 'document', 'prospektus'),
(9, 'Prospektus Tomat Bogor', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/documents/Dokumen-Investasi-Tomat-Bogor.pdf', 'document', 'prospektus');

INSERT INTO project_files (project_id, filename, path, label, purpose) VALUES
(1, 'CKB18081-Lahan Mitigasi Sawi putih', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/CKB18081-Lahan+Mitigasi+Sawi+putih.jpg', 'image', 'gallery'),
(1, 'CKB18081-Lahan Mitigasi Sawi putih', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/CKB18081-Lahan+Mitigasi+Sawi+putih2.jpg', 'image', 'gallery'),
(1, 'CKB18081-Lahan Selesai Pemasangan Mulsa', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/CKB18081-Lahan+Selesai+Pemasangan+Mulsa.jpg', 'image', 'gallery'),
(1, 'CKB18081-Lahan Selesai Pemasangan Mulsa', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/CKB18081-Lahan+Selesai+Pemasangan+Mulsa2.jpg', 'image', 'gallery'),
(1, 'CKBGT - Kondisi Tanaman', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/CKBGT+-+Kondisi+Tanaman+(1).jpg', 'image', 'gallery'),
(1, 'CKBGT - Pemasangan Mulsa', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/CKBGT+-+Pemasangan+Mulsa+(4).jpg', 'image', 'gallery'),
(1, 'CKBGT - Pembuatan Bedengan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/CKBGT+-+Pembuatan+Bedengan+(4).jpg', 'image', 'gallery'),
(1, 'CKBGT - Pengecoran Pupuk', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/CKBGT+-+Pengecoran+Pupuk+(2).jpg', 'image', 'gallery'),
(1, 'CKBGT - Penyemprotan Pestisida', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/CKBGT+-+Penyemprotan+Pestisida+(1).jpg', 'image', 'gallery'),
(1, 'CKBGT - Penyiangan Gulma', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/CKBGT+-+Penyiangan+Gulma+(6).jpg', 'image', 'gallery'),
(1, 'CKBGT - Penyiangan Gulma', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/CKBGT+-+Penyiangan+Gulma+(7).jpg', 'image', 'gallery'),
(1, 'CKBGT - Pemanenan Cabai Kriting', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/Pemanenan+Cabe+Keriting+(1)(1).jpg', 'image', 'gallery'),
(1, 'CKBGT - Pemanenan Cabai Kriting', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/Pemanenan+Cabe+Keriting+(1)(1).jpg', 'image', 'gallery'),
(1, 'CKBGT - Pemanenan Cabai Kriting', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/Pemanenan+Cabe+Keriting+(3).jpg', 'image', 'gallery'),
(1, 'CKBGT - Pemanenan Cabai Kriting', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/Pemanenan+Cabe+Keriting+(4)(1).jpg', 'image', 'gallery'),
(1, 'CKBGT - Pemanenan Cabai Kriting', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-keriting/Pemanenan+Cabe+Keriting+(4).jpg', 'image', 'gallery');

INSERT INTO project_files (project_id, filename, path, label, purpose) VALUES
(2, 'BB18051 - Brokoli yang siap panen', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/brokoli/BB18051+-+Brokoli+yang+siap+panen.jpg', 'image', 'gallery'),
(2, 'BB18051 - Packing hasil panen brokoli', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/brokoli/BB18051+-+Packing+hasil+panen+Brokoli.jpg', 'image', 'gallery'),
(2, 'BB18051 - Pemanenan Brokoli', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/brokoli/BB18051+-+Pemanenan+Brokoli.jpg', 'image', 'gallery'),
(2, 'BB18051 - Pengamatan kondisi tanaman', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/brokoli/BB18051+-+Pengamatan+kondisi+tanaman.jpg', 'image', 'gallery'),
(2, 'BB18051 - Pengankutan hasil panen', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/brokoli/BB18051+-+Pengangkutan+hasil+panen.jpg', 'image', 'gallery'),
(2, 'BB18051 - Persemaian brokoli untuk penyulaman', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/brokoli/BB18051+-+Persemaian+Brokoli+untuk+Penyulaman+(1).jpg', 'image', 'gallery'),
(2, 'Melakukan pemilihan kualitas brokoli', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/brokoli/Melakukan+Pemilihan+Kualitas+Brokoli(1).jpg', 'image', 'gallery'),
(2, 'Melakukan pemilihan kualitas brokoli', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/brokoli/Melakukan+Pemilihan+Kualitas+Brokoli.jpg', 'image', 'gallery'),
(2, 'Pengemasan brokoli', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/brokoli/Pengemasan+Brokoli.jpg', 'image', 'gallery'),
(2, 'Panen', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/brokoli/Panen.jpg', 'image', 'gallery');

INSERT INTO project_files (project_id, filename, path, label, purpose) VALUES
(3, 'DBB-GT - Panen', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/daun-bawang/DBB-GT+-+Panen+(1).jpg', 'image', 'gallery'),
(3, 'DBB-GT - Panen', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/daun-bawang/DBB-GT+-+Panen+(2).jpg', 'image', 'gallery'),
(3, 'DBB18051 - Kondisi tanaman daun bawang', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/daun-bawang/DBB18051+-+Kondisi+Tanaman+Daun+Bawang+(1).jpg', 'image', 'gallery'),
(3, 'DBB18051 - Kondisi tanaman daun bawang', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/daun-bawang/DBB18051+-+Kondisi+Tanaman+Daun+Bawang+(2).jpg', 'image', 'gallery'),
(3, 'DBB18053 - Pemberian pupuk kandang', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/daun-bawang/DBB18053+-+Pemberian+Pupuk+Kandang+(1).jpg', 'image', 'gallery'),
(3, 'DBB18053 - Pemberian pupuk kandang', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/daun-bawang/DBB18053+-+Pemberian+Pupuk+Kandang+(8).jpg', 'image', 'gallery'),
(3, 'Pengendalian hama dan penyakit', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/daun-bawang/Pengendalian+Hama+dan+Penyakit.jpg', 'image', 'gallery'),
(3, 'Pengendalian hama dan penyakit', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/daun-bawang/Pengendalian+Hama+dan+Penyakit(1).jpg', 'image', 'gallery');

INSERT INTO project_files (project_id, filename, path, label, purpose) VALUES
(4, 'KKB18051 - Pemanenan Kacang Kapri', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kacang-kapri/KKB18051+-+Pemanenan+Kacang+Kapri+(2).jpg', 'image', 'gallery'),
(4, 'KKB18051 - Pemanenan Kacang Kapri', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kacang-kapri/KKB18051+-+Pemanenan+Kacang+Kapri.jpg', 'image', 'gallery'),
(4, 'KKB18051 - Pengolahan lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kacang-kapri/KKB18051+-+Pengolahan+lahan+(2).jpg', 'image', 'gallery'),
(4, 'KKB18051 - Pengolahan lahan menggunakan kerbau', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kacang-kapri/KKB18051+-+Pengolahan+lahan+menggunakan+kerbau+(1).jpg', 'image', 'gallery'),
(4, 'KKB18062 - Kacang kapri yang siap panen', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kacang-kapri/KKB18062+-+Kacang+Kapri+yang+siap+panen.jpg', 'image', 'gallery'),
(4, 'KKB18062 - Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kacang-kapri/KKB18062+-+Kondisi+lahan+(2).jpg', 'image', 'gallery'),
(4, 'KKB18062 - Panen kacang kapri', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kacang-kapri/KKB18062+-+Panen+Kacang+Kapri+(1).jpg', 'image', 'gallery'),
(4, 'KKB18062 - Panen kacang kapri', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kacang-kapri/KKB18062+-+Panen+Kacang+Kapri+(2).jpg', 'image', 'gallery'),
(4, 'KKB18062 - Tanaman kacang kapri mulai berbunga', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kacang-kapri/KKB18062+-+Tanaman+Kacang+Kapri+mulai+berbunga.jpg', 'image', 'gallery');

INSERT INTO project_files (project_id, filename, path, label, purpose) VALUES
(5, 'PCB18051 - Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/pak-choy/PCB18051+-+Kondisi+lahan+(1).jpg', 'image', 'gallery'),
(5, 'PCB18051 - Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/pak-choy/PCB18051+-+Kondisi+tanaman+Pak+Choy.jpg', 'image', 'gallery'),
(5, 'PCB18052 - Persemaian pak choy', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/pak-choy/PCB18052+-+Persemaian+Pak+Choy+(1).jpg', 'image', 'gallery'),
(5, 'Pak choy siap panen', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/pak-choy/Pakchoy+siap+panen.jpg', 'image', 'gallery'),
(5, 'Panen pak choy', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/pak-choy/Panen+Pak+Choy.jpeg', 'image', 'gallery'),
(5, 'Panen pak choy', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/pak-choy/Panen+Pak+Choy.jpg', 'image', 'gallery'),
(5, 'Pengemasan hasil panen pak choy', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/pak-choy/Pengemasan+Hasil+Panen+Pak+Choy.jpg', 'image', 'gallery'),
(5, 'Pengemasan pak choy', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/pak-choy/Pengemasan+Pak+Choy.jpg', 'image', 'gallery');

INSERT INTO project_files (project_id, filename, path, label, purpose) VALUES
(6, 'Panen sawi putih', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/sawi-putih/Panen+Sawi+Putih.jpg', 'image', 'gallery'),
(6, 'Pemilihan sawi berkualitas', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/sawi-putih/Pemilihan+Sawi+Berkualitas.jpg', 'image', 'gallery'),
(6, 'Pengangkutan pemanenan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/sawi-putih/Pengankutan+Pemanenan.jpg', 'image', 'gallery'),
(6, 'Pemberian pupuk NPK phonska', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/sawi-putih/SPB18052+-+Pemberian+Pupuk+NPK+Phonska+(2).jpg', 'image', 'gallery'),
(6, 'SPB18052 - Pembuatan bedengan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/sawi-putih/SPB18052+-+Pembuatan+Bedengan+(1).jpg', 'image', 'gallery'),
(6, 'SPB18053 - Packing sawi putih', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/sawi-putih/SPB18053-Packing+Sawi+Putih.jpg', 'image', 'gallery'),
(6, 'SPB18053 - Pemanenan sawi putih', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/sawi-putih/SPB18053-Pemanenan+Sawi+Putih+(2).jpg', 'image', 'gallery'),
(6, 'SPB18053 - Pengangkutan hasil panen', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/sawi-putih/SPB18053-Pengangkutan+hasil+panen+(1).jpg', 'image', 'gallery'),
(6, 'SPB18053 - Pengangkutan hasil panen', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/sawi-putih/SPB18053-Pengangkutan+hasil+panen+(2).jpg', 'image', 'gallery');

INSERT INTO project_files (project_id, filename, path, label, purpose) VALUES
(7, 'Cabai rawit mulai menguning', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-rawit/Cabai+rawit+mulai+menguning.jpg', 'image', 'gallery'),
(7, 'Proyek cabai rawit', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-rawit/IMG_20180604_153521.jpg', 'image', 'gallery'),
(7, 'Kondisi lahan cabai rawit', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-rawit/Kondisi+Lahan+Cabe+Rawit.jpg', 'image', 'gallery'),
(7, 'Lahan cabai rawit', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-rawit/Lahan+Cabai+rawit.jpg', 'image', 'gallery'),
(7, 'Lahan cabai rawit', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-rawit/Lahan+Cabai+rawit(1).jpg', 'image', 'gallery'),
(7, 'Lahan cabai rawit', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-rawit/Lahan+Cabe+rawit.jpg', 'image', 'gallery'),
(7, 'Panen cabai rawit', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-rawit/Panen+Cabe+Rawit(1).jpg', 'image', 'gallery'),
(7, 'Panen cabai rawit', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/cabai-rawit/Panen+Cabe+Rawit.jpg', 'image', 'gallery');

INSERT INTO project_files (project_id, filename, path, label, purpose) VALUES
(8, 'Kondisi Lahan Kentang Granola di Berastagi', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kentang-granola/Kondisi+Lahan+Kentang+Granola+di+Brastagi.jpg', 'image', 'gallery'),
(8, 'Pemasangan guludan sebagai drainase', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kentang-granola/Pemasangan+guludan+sebagai+drainase+(1).jpg', 'image', 'gallery'),
(8, 'Pemasangan guludan sebagai drainase', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kentang-granola/Pemasangan+guludan+sebagai+drainase+(2).jpg', 'image', 'gallery'),
(8, 'Pemasangan guludan sebagai drainase', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kentang-granola/Pemasangan+guludan+sebagai+drainase+(3).jpg', 'image', 'gallery'),
(8, 'Pembukaan lahan kentang granola', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kentang-granola/Pembukaan+Lahan+Kentang+Granola.jpg', 'image', 'gallery'),
(8, 'Penanaman kentang granola di Berastagi', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kentang-granola/Penanaman+Kentang+Granola+di+Brastagi(1).jpg', 'image', 'gallery'),
(8, 'Penanaman kentang granola di Berastagi', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kentang-granola/Penanaman+Kentang+Granola+di+Brastagi.jpg', 'image', 'gallery'),
(8, 'Tanaman kentang granola', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/kentang-granola/Tanaman+Kentang+Granola.jpg', 'image', 'gallery');

INSERT INTO project_files (project_id, filename, path, label, purpose) VALUES
(9, 'Hamparan lahan tomat mang peya siap panen', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/Hamparan+lahan+tomat+Mang+Peya+siap+panen.jpg', 'image', 'gallery'),
(9, 'Panen tomat', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/Panen+Tomat.jpg', 'image', 'gallery'),
(9, 'Panen tomat di kebung mang peya', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/Panen+tomat+di+kebun+Mang+Peya.jpg', 'image', 'gallery'),
(9, 'Kondisi bedengan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB18101+-+Kondisi+bedengan.jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB18101+-+Kondisi+lahan+(1).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB18101+-+Kondisi+lahan+(2).jpg', 'image', 'gallery'),
(9, 'Pengolahan lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB18101+-+Pengolahan+lahan.jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181020+-+Kondisi+lahan+(1).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181020+-+Kondisi+lahan+(2).jpg', 'image', 'gallery'),
(9, 'Pengolahan lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181020+-+Pengolahan+lahan.jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181025+-+Kondisi+lahan+(1).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181025+-+Kondisi+lahan+(2).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181025+-+Kondisi+lahan+(3).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181025+-+Kondisi+lahan+(4).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181027+-+Kondisi+lahan+(1).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181027+-+Kondisi+lahan+(2).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181027+-+Kondisi+lahan+(3).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181027+-+Kondisi+lahan+(4).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181028+-+Kondisi+lahan+(1).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181028+-+Kondisi+lahan+(2).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181028+-+Kondisi+lahan+(3).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181028+-+Kondisi+lahan+(4).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181038+-+Kondisi+lahan+(1).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181038+-+Kondisi+lahan+(2).jpg', 'image', 'gallery'),
(9, 'Kondisi lahan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181038+-+Kondisi+lahan+(3).jpg', 'image', 'gallery'),
(9, 'Survey lahan bersama petani', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/TB181038+-+Survey+lahan+bersama+petani.jpg', 'image', 'gallery'),
(9, 'Tomat umur 2 bulan', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/gallery/tomat-bogor/Tomat+umur+2+bulan.jpg', 'image', 'gallery');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

