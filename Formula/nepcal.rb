class Nepcal < Formula
  desc "Equivalent of Linux's cal, for Nepali dates"
  homepage "https://github.com/nepcal/nepcal"
  url "https://github.com/srishanbhattarai/nepcal/archive/v0.3.0.tar.gz"
  head "https://github.com/srishanbhattarai/nepcal.git"
  version "0.3.0"
  sha256 "e5033a604c00254779a93cf3194d27c82f30e428b131fcd2b9127b13ac484d71"
  depends_on "make" => :build

  def install
    bin.install "nepcal"
  end
end
